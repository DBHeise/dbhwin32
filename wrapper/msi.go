package wrapper

import (
	"syscall"
	"unsafe"
)

var (
	modmsi = syscall.NewLazyDLL("msi.dll")

	procMsiEnumClients    = modmsi.NewProc("MsiEnumClientsW")
	procMsiEnumProducts   = modmsi.NewProc("MsiEnumProductsW")
	procMsiEnumPatches    = modmsi.NewProc("MsiEnumPatchesW")
	procMsiEnumComponents = modmsi.NewProc("MsiEnumComponentsW")

	procMsiGetComponentPath = modmsi.NewProc("MsiGetComponentPathW")
)

func handleEnumReturn(r uintptr) error {
	err := syscall.Errno(r)
	if err != syscall.Errno(0) {
		return err
	}
	return nil
}

func MsiEnumProducts(index uint32, buffer *uint16) error {
	//https://msdn.microsoft.com/en-us/library/windows/desktop/aa370101(v=vs.85).aspx
	r, _, _ := syscall.Syscall(
		procMsiEnumProducts.Addr(),
		2,
		uintptr(index),
		uintptr(unsafe.Pointer(buffer)),
		0)
	return handleEnumReturn(r)
}

func MsiEnumComponents(index uint32, buffer *uint16) error {
	//https://msdn.microsoft.com/en-us/library/windows/desktop/aa370097(v=vs.85).aspx
	r, _, _ := syscall.Syscall(
		procMsiEnumComponents.Addr(),
		2,
		uintptr(index),
		uintptr(unsafe.Pointer(buffer)),
		0)
	return handleEnumReturn(r)
}

func MsiEnumClients(componentGUID *uint16, index uint32, buffer *uint16) error {
	//https://msdn.microsoft.com/en-us/library/windows/desktop/aa370094(v=vs.85).aspx
	r, _, _ := syscall.Syscall(
		procMsiEnumClients.Addr(),
		3,
		uintptr(unsafe.Pointer(componentGUID)),
		uintptr(index),
		uintptr(unsafe.Pointer(buffer)))
	return handleEnumReturn(r)
}

func MsiEnumPatches(productGUID *uint16, index uint32, buffer *uint16, transformBuffer *uint16, transformBufferSize *uint32) error {
	r, _, _ := syscall.Syscall6(
		procMsiEnumPatches.Addr(),
		5,
		uintptr(unsafe.Pointer(productGUID)),
		uintptr(index),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(transformBuffer)),
		uintptr(unsafe.Pointer(transformBufferSize)),
		0)
	return handleEnumReturn(r)
}

func MsiGetComponentPath(productGUID *uint16, componentGUID *uint16, buffer *uint16, bufferSize *uint32) error {
	//https://msdn.microsoft.com/en-us/library/windows/desktop/aa370112(v=vs.85).aspx
	r, _, _ := syscall.Syscall6(
		procMsiGetComponentPath.Addr(),
		4,
		uintptr(unsafe.Pointer(productGUID)),
		uintptr(unsafe.Pointer(componentGUID)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(bufferSize)),
		0,
		0)
	return handleEnumReturn(r)
}
