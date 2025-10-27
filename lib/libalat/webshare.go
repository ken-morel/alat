package main

import "C"
import (
	"fmt"
	"strings"
)

//export wehsbare_get_status_json
func wehsbare_get_status_json(handle C.int) *C.char {
	if instance := getInstance(handle); instance != nil {
		return toJSON(*instance.node.Services.WebShare.GetStatus())
	} else {
		return nil
	}
}

//export webshare_add_shared_files
func webshare_add_shared_files(handle C.int, filesC *C.char) C.int {
	if instance := getInstance(handle); instance != nil {
		filePathsCombo := C.GoString(filesC)
		files := strings.Split(filePathsCombo, ";")
		err := instance.node.Services.WebShare.AddSharedFiles(files)
		if err != nil {
			setError(err)
			return -2
		} else {
			return 0
		}
	} else {
		return -1
	}
}

//export webshare_remove_shared_files_by_uuids
func webshare_remove_shared_files_by_uuids(handle C.int, filesC *C.char) C.int {
	if instance := getInstance(handle); instance != nil {
		uuids := strings.Split(C.GoString(filesC), ";")
		instance.node.Services.WebShare.RemoveSharedFilesByUUIDS(uuids)
		return 0
	} else {
		return -1
	}
}

//export webshare_remove_shared_files_by_paths
func webshare_remove_shared_files_by_paths(handle C.int, filesC *C.char) C.int {
	if instance := getInstance(handle); instance != nil {
		paths := strings.Split(C.GoString(filesC), ";")
		instance.node.Services.WebShare.RemoveSharedFilesByPaths(paths)
		return 0
	} else {
		noSuchInstance(handle)
		return -1
	}
}

//export webshare_clear_shared_files
func webshare_clear_shared_files(handle C.int) C.int {
	if instance := getInstance(handle); instance != nil {
		instance.node.Services.WebShare.ClearSharedFiles()
		return 0
	} else {
		return -1
	}
}

//export webshare_get_shared_files_json
func webshare_get_shared_files_json(handle C.int) *C.char {
	if instance := getInstance(handle); instance != nil {
		return toJSON(instance.node.Services.WebShare.GetSharedFiles())
	} else {
		return nil
	}
}

//export webshare_get_passcode
func webshare_get_passcode(handle C.int) *C.char {
	if instance := getInstance(handle); instance != nil {
		return C.CString(instance.node.Services.WebShare.GetPasscode())
	} else {
		return nil
	}
}

//export webshare_set_passcode
func webshare_set_passcode(handle C.int, passcodeC *C.char) C.int {
	if instance := getInstance(handle); instance != nil {
		instance.node.Services.WebShare.SetPasscode(C.GoString(passcodeC))
		return 0
	} else {
		return -1
	}
}

//export webshare_start
func webshare_start(handle C.int) C.int {
	if instance := getInstance(handle); instance != nil {
		port, err := instance.node.Services.WebShare.Start()
		if err != nil {
			setError(fmt.Errorf("failed to start webshare service: %v", err))
			return -2
		} else {
			return C.int(port)
		}
	} else {
		return -1
	}
}

//export webshare_stop
func webshare_stop(handle C.int) C.int {
	if instance := getInstance(handle); instance != nil {
		err := instance.node.Services.WebShare.Stop()
		if err != nil {
			setError(fmt.Errorf("failed to stop webshare service: %v", err))
			return -2
		} else {
			return 0
		}
	} else {
		return -1
	}
}
