package transmit

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"
)

var (
	DefaultSShTcpTimeout = 15 * time.Second
	InvalidHostName = errors.New("invalid parameters: hostname is empty")
	InvalidPort = errors.New("invalid parameters: port must be range 0 ~ 65535")
)

type SSHConfig struct {
	Ip			string			`ini:"ip"`
	Port		uint32			`ini:"port"`
	User 		string			`ini:"user"`
	Password 	string			`ini:"password"`
	KeyFile 	string			`ini:"key_file"`
	Timeout 	time.Duration	`ini:"timeout"`
}

type AuthConfig struct {
	*ssh.ClientConfig
	*SSHConfig
}

type SSHClient struct {
	client *ssh.Client
	AuthConfig AuthConfig
}

// ExecInfo 存放执行结果的结构体信息
type ExecInfo struct {
	Cmd         string
	Output     	[]byte
	ExitCode 	int
}

func (e *ExecInfo) String() string {
	return fmt.Sprintf(`ExecInfo(cmd: "%s", exitcode: %d)`, e.Cmd, e.ExitCode)
}

func (e *ExecInfo) OutputString() string {
	return string(e.Output)
}

func (a *AuthConfig) SetAuthMethod() (ssh.AuthMethod, error) {
	a.setDefault()

	if a.Password != "" {
		return ssh.Password(a.Password), nil
	}

	data, err := ioutil.ReadFile(a.KeyFile)

	if err != nil {
		return nil, err
	}

	singer, err := ssh.ParsePrivateKey(data)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(singer), nil
}

func (a *AuthConfig) ApplyConfig() error {
	authMethod, err := a.SetAuthMethod()
	if err != nil {
		return err
	}

	a.ClientConfig = &ssh.ClientConfig{
		User: a.SSHConfig.User,
		Auth: []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: a.SSHConfig.Timeout,
	}

	return nil
}

// Connect 与远程主机连接
func (s *SSHClient) Connect() error {
	if s.client != nil {
		log.Println("Already Login")
		return nil
	}

	if err := s.AuthConfig.ApplyConfig(); err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", s.AuthConfig.Ip, s.AuthConfig.Port)

	var err error

	s.client, err = ssh.Dial("tcp", addr, s.AuthConfig.ClientConfig)

	if err != nil {
		return err
	}

	return nil
}

// NewSSHClient SSHclient的构造方法
func NewSSHClient(sshConfig *SSHConfig) (*SSHClient, error) {
	authConfig := AuthConfig{nil, sshConfig}

	switch {
	case authConfig.SSHConfig.Ip == "":
		return nil, InvalidHostName
	case authConfig.SSHConfig.Port > 65535 || authConfig.SSHConfig.Port < 0:
		return nil, InvalidPort
	}

	sshClient := &SSHClient{AuthConfig: authConfig}
	err := sshClient.Connect()
	if err != nil {
		return nil, err
	}

	return sshClient, nil
}

// Exec 一个session只能执行一次命令，也就是说不能在同一个session执行多次s.session.CombinedOutput
//如果想执行多次，需要每条为每个命令创建一个session(这里是这样做)
func (s *SSHClient) Exec(cmd string) (*ExecInfo, error) {
	session, err := s.client.NewSession()

	if err != nil {
		return nil, err
	}

	defer session.Close()

	output, err := session.CombinedOutput(cmd)

	var exitcode int
	if err != nil {
		// 断言转成具体实现类型，获取返回值
		exitcode = err.(*ssh.ExitError).ExitStatus()
	}

	return &ExecInfo {
		Cmd: cmd,
		Output: output,
		ExitCode: exitcode,
	}, nil
}

// 返回当前用户名
func getCurrentUser() string {
	user, _ := user.Current()
	return user.Username
}

func (a *AuthConfig) setDefault() {
	if a.SSHConfig.User == "" {
		a.SSHConfig.User = getCurrentUser()
	}

	if a.KeyFile == "" {
		userHome, _ := os.UserHomeDir()
		a.KeyFile = fmt.Sprintf("%s/.ssh/id_rsa", userHome)
	}

	if a.SSHConfig.Timeout == 0 {
		a.SSHConfig.Timeout = DefaultSShTcpTimeout
	}
}




