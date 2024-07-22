// Code generated by command: go run mask_asm.go -pkg mem -out ../mem/mask_amd64.s -stubs ../mem/mask_amd64.go. DO NOT EDIT.

//go:build !purego

package mem

// Mask set bits of dst to zero and copies the one-bits of src to dst, returning the number of bytes written.
func Mask(dst []byte, src []byte) int
