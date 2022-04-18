package remind_me

import (
	"github.com/Logiase/MiraiGo-Template/bot"
	"sync"
)

func init()  {
	instance = &reminder{}                           //实例化一个结构体对象
	bot.RegisterModule(instance)
}

type reminder struct {
}


func (m *reminder) MiraiGoModule() bot.ModuleInfo {         //skl（struct）绑定了一个方法
	return bot.ModuleInfo{
		ID:       "t0hka.reminder",
		Instance: instance,
	}
}

func (m *reminder) Init()  {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
}

func (m *reminder) PostInit() {
	// 第二次初始化
	// 再次过程中可以进行跨Module的动作
	// 如通用数据库等等
}

func (m *reminder) Serve(b *bot.Bot) {
	// 注册服务函数部分
	b.OnPrivateMessage(Reminder)

}

func (m *reminder) Start(b *bot.Bot) {
	// 此函数会新开携程进行调用
	// ```go
	// 		go exampleModule.Start()
	// ```

	// 可以利用此部分进行后台操作
	// 如http服务器等等
}
func (m *reminder) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
}

var instance *reminder