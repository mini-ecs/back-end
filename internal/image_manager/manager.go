package image_manager

import (
	"github.com/mini-ecs/back-end/pkg/log"
	"io"
	"os"
)

// https://github.com/willscott/go-nfs-client
// https://github.com/willscott/go-nfs

type ImageManager interface {
	Copy(new, old string) error
	Delete(image string) error
}

var LocalMachineImpl ImageManager = &localMachineImpl{}

type localMachineImpl struct {
}

func (l *localMachineImpl) Copy(new, old string) error {
	log.GetGlobalLogger().Infof("destination=%v, source=%v", new, old)
	src, err := os.Open(old)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(new)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func (l *localMachineImpl) Delete(image string) error {
	log.GetGlobalLogger().Infof("Deleteing image: %v", image)
	err := os.Remove(image)
	if err != nil {
		return err
	}
	return nil
}
