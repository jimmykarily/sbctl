package main

import (
	"testing"

	"github.com/foxboron/go-uefi/efi"
	"github.com/foxboron/sbctl/tests/utils"
	"github.com/hugelgupf/vmtest/guest"
)

func TestEnrollKeys(t *testing.T) {
	guest.SkipIfNotInVM(t)

	if efi.GetSecureBoot() {
		t.Fatal("in secure boot mode")
	}

	if !efi.GetSetupMode() {
		t.Fatal("not in setup mode")
	}

	utils.Exec("rm -rf /usr/share/secureboot")
	utils.Exec("sbctl status")
	utils.Exec("sbctl create-keys")
	if out, err := utils.ExecWithOutput("sbctl enroll-keys"); err == nil {
		t.Fatalf("Expected error about \"Could not find any TPM Eventlog in the system\", none happened. Command output: %s", out)
	}

	if out, err := utils.ExecWithOutput("sbctl enroll-keys --yes-this-might-brick-my-machine"); err != nil {
		t.Fatalf("%s: %s", err.Error(), out)
	}

	if efi.GetSetupMode() {
		t.Fatal("in setup mode")
	}

}
