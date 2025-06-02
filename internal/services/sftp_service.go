package services

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	sftpUser     = "demo"
	sftpPassword = "password"
	sftpHost     = "test.rebex.net"
	sftpPort     = "22"
	sftpFilePath = "/pub/example/readme.txt"  
)

func DownloadAndProcessFile() error {
	config := &ssh.ClientConfig{
		User: sftpUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sftpPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", sftpHost, sftpPort), config)
	if err != nil {
		return fmt.Errorf("SFTP dial error: %v", err)
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		return fmt.Errorf("SFTP client error: %v", err)
	}
	defer client.Close()

	remoteFile, err := client.Open(sftpFilePath)
	if err != nil {
		return fmt.Errorf("error opening remote file: %v", err)
	}
	defer remoteFile.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, remoteFile); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	log.Println("✅ File content:")
	log.Println(buf.String())

	return nil
}
