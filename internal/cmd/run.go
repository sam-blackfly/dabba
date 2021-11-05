package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/sam-blackfly/dabba/internal/paths"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute command inside a container",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Forks command inside a container",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fork(args)
	},
}

func run(args []string) {
	cmd := exec.Command("/proc/self/exe", append([]string{"fork"}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		Credential: &syscall.Credential{Uid: 0, Gid: 0},
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getgid(), Size: 1},
		},
	}

	ensure(cmd.Run())
}

func fork(args []string) {
	log.Printf("Running %v as pid %v\n", args, os.Getpid())

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	chrootPath := path.Join(paths.FileSystemsPath, "alpine")

	ensure(syscall.Sethostname([]byte("dabba")))
	ensure(syscall.Chroot(chrootPath))
	ensure(syscall.Chdir("/"))
	ensure(syscall.Mount("proc", "proc", "proc", 0, ""))

	ensure(cmd.Run())

	ensure(syscall.Unmount("proc", 0))
}

func ensure(err error) {
	if err != nil {
		panic(err)
	}
}
