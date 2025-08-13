package golib

import "runtime"

// Version retorna a versão do Go (note o V maiúsculo para exportar)
func Version() string {
    return runtime.Version()
}
