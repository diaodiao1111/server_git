package Interface

import "fmt"

type In interface {
	start()
	stop()
}

type Phone struct {
}

type Camara struct {
}

func (p Phone) start() {
	fmt.Println("手机开始")
}

func (p Phone) stop() {
	fmt.Println("手机关闭")
}

func (c Camara) start() {
	fmt.Println("相机开始")
}

func (c Camara) stop() {
	fmt.Println("相机开始")
}

func Computor(in In) {

	in.start()
	in.stop()

}
