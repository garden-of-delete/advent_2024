package main

import "fmt"

type Pos struct {
	x int
	y int
}

type Direction struct {
	xDir int
	yDir int
}

type VisitationRecord struct {
	Pos
	Direction
}

var (
	UP    = Direction{0, -1}
	DOWN  = Direction{0, 1}
	LEFT  = Direction{-1, 0}
	RIGHT = Direction{1, 0}
)

type World struct {
	guardPos    Pos
	guardDir    Direction
	obstaclePos *Set[Pos]
	xSize       int
	ySize       int
	visited     map[Pos][]Direction // stores the orientation of the guard when exiting a Pos
}

func (world *World) turnRight() {

	if world.guardDir == UP {
		world.guardDir = RIGHT
	} else if world.guardDir == RIGHT {
		world.guardDir = DOWN
	} else if world.guardDir == DOWN {
		world.guardDir = LEFT
	} else { // world.guardDir == LEFT
		world.guardDir = UP
	}
}

func (world *World) collisionCheck() bool {

	if world.obstaclePos.Contains(world.guardPos) {
		return true
	}
	return false
}

func (world *World) checkWorldExit() bool {

	if world.guardPos.x < 0 ||
		world.guardPos.y < 0 ||
		world.guardPos.x == world.xSize ||
		world.guardPos.y == world.ySize {
		return true
	}
	return false
}

func (world *World) addObstacle(pos Pos) {

	if world.guardPos == pos { // don't add an obstacle at the initial position
		return
	}
	world.obstaclePos.Add(pos)
}

func (world *World) removeObstacle(pos Pos) {

	world.obstaclePos.Remove(pos)
}

func (world *World) printWorld() {

	for y := 0; y < world.ySize; y++ {
		for x := 0; x < world.xSize; x++ {
			if x == world.guardPos.x && y == world.guardPos.y {
				if world.guardDir == UP {
					print("^")
				} else if world.guardDir == DOWN {
					print("v")
				} else if world.guardDir == LEFT {
					print("<")
				} else {
					print(">")
				}
			} else if world.obstaclePos.Contains(Pos{x, y}) {
				print("#")
			} else if _, exists := world.visited[Pos{x, y}]; exists {
				print("X")
			} else {
				print(".")
			}
		}
		print("\n")
	}
}

func NewWorld(lines []string) *World {

	world := World{}
	world.obstaclePos = NewSet[Pos]()
	world.visited = map[Pos][]Direction{}
	world.ySize = len(lines)
	world.xSize = len(lines[0])
	for y := range lines {
		chars := []rune(lines[y])
		for x := range chars {
			if chars[x] == '.' {
				continue
			} else if chars[x] == '#' {
				world.obstaclePos.Add(Pos{x, y})
			} else if chars[x] == '^' {
				world.guardPos.x = x
				world.guardPos.y = y
				world.guardDir = UP
			} else {
				fmt.Println("ERROR: invalid input character: ", chars[x])
			}
		}
	}
	return &world
}

func (world *World) checkCycle() bool {

	if hasDuplicates(world.visited[world.guardPos]) {
		return true
	}
	return false
}

func (world *World) Step() bool { // no invalid worlds allowed outside this func

	currentPos := world.guardPos
	currentDir := world.guardDir
	peekPos := Pos{
		world.guardPos.x + world.guardDir.xDir,
		world.guardPos.y + world.guardDir.yDir,
	}
	world.guardPos = peekPos
	if world.checkWorldExit() {
		world.visited[currentPos] = append(world.visited[currentPos], currentDir)
		return false
	} else if world.collisionCheck() {
		world.guardPos = currentPos
		world.turnRight()
		return world.Step()
	} else {
		world.visited[currentPos] = append(world.visited[currentPos], currentDir)
		return true // peekPos becomes currentPos
	}
}

func runWorldSim(world *World) bool {

	for nSteps := 1; world.Step(); nSteps++ {
		if world.checkCycle() {
			return true // world contains a cycle
		}
	}
	return false
}

func daySix() {

	lines := fileLineScanner("input-data/day6_input.txt")

	initWorld := NewWorld(lines)
	initWorld.printWorld()

	// initial run
	runWorldSim(initWorld)
	fmt.Printf("INFO: guard visited %d distinct locations\n", len(initWorld.visited))

	// for every position on the initial path
	nCycles := 0
	for pos := range initWorld.visited {
		world := NewWorld(lines)
		world.addObstacle(pos)
		if runWorldSim(world) {
			nCycles++
		}
	}
	fmt.Printf("INFO: %d possible cycles\n", nCycles)
}
