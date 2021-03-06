package commands

import (
	"os"
	"os/exec"

	// "k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
)

func (c *Command) Upgrade() (*services.UpdateReleaseResponse, error) {
	ns := "default"
	cmd := exec.Command("helm", "upgrade", c.Release, c.Chart)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, flag := range c.Flags {
		for k, v := range flag {
			switch k {
			// Local
			case "disable-hooks":
				cmd.Args = append(cmd.Args, "--disable-hooks")
			case "dry-run":
				cmd.Args = append(cmd.Args, "--dry-run")
			case "i":
				fallthrough
			case "install":
				cmd.Args = append(cmd.Args, "--install")
			case "keyring":
				cmd.Args = append(cmd.Args, "--keyring", v.(string))
			case "namespace":
				ns = v.(string)
			case "set":
				cmd.Args = append(cmd.Args, "--set", v.(string))
			case "f":
				fallthrough
			case "values":
				cmd.Args = append(cmd.Args, "--values", v.(string))
			case "verify":
				cmd.Args = append(cmd.Args, "--verify")
			case "version":
				cmd.Args = append(cmd.Args, "--version", v.(string))

			// Global
			case "debug":
				cmd.Args = append(cmd.Args, "--debug")
			case "home":
				cmd.Args = append(cmd.Args, "--home", v.(string))
			case "host":
				cmd.Args = append(cmd.Args, "--host", v.(string))
			case "kube-context":
				cmd.Args = append(cmd.Args, "--kube-context", v.(string))
			}
		}
	}

	cmd.Args = append(cmd.Args, "--namespace", ns)
	trace(cmd)

	err := cmd.Run()
	if err != nil {
		c.Args = append([]string{}, "0")
		c.Rollback()
	}

	return nil, err
}
