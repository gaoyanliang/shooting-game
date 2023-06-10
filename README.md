# shooting-game

一款使用 Ebiten 开发的简易射击游戏

![game](./images/game.png)

- ebiten：ebiten 引擎的一些基础用法
- images：游戏中用到的一些图片
- bullet：子弹
- config：游戏配置
- game：游戏主逻辑
- monster：怪物👹
- main：启动

## TODO

1. [x] 新增飞船
2. [x] 新增怪物
3. [x] 新增子弹 
4. [x] 当前屏幕中只允许出现固定数量的子弹，超出不允许再射击 
5. [x] 新增子弹击打怪物逻辑
6. [x] 新增游戏开始/结束界面
7. [x] 新增判断游戏胜负逻辑
8. [x] 新增记分板

## 如何体验？

### 1️⃣ 克隆这个 repo，进入代码文件夹

```shell
git clone https://github.com/gaoyanliang/shooting-game.git
cd shooting-game
```

### 2️⃣ 安装依赖项

```shell
go mod download
```

### 3️⃣ 编译生成可执行文件 shooting-game

```shell
go build -o shooting-game *.go
```

### 4⃣️ 运行可执行文件 shooting-game

```shell
./shooting-game
```

## 配套教程

- [Ebitengine - 一款Go语言编写的2D游戏引擎](https://yanliang.cool/blog/go/ebitengine/)
- [用 Ebitengine 制作一款简易射击游戏](https://yanliang.cool/blog/go/shooting-game/)