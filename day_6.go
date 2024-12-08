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
	visited     map[Pos]Direction
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
		world.visited[currentPos] = currentDir
		return false
	} else if world.collisionCheck() {
		world.guardPos = currentPos
		world.turnRight()
		fmt.Println("INFO: Turning right...")
		return world.Step()
	} else {
		world.visited[currentPos] = currentDir
		return true // peekPos becomes currentPos
	}
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
	world.visited = map[Pos]Direction{}
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
				world.visited[Pos{x, y}] = UP
			} else {
				fmt.Println("ERROR: invalid input character: ", chars[x])
			}
		}
	}
	return &world
}

func daySix() {

	// read initial world
	//lines := fileLineScanner("input-data-test/day6_input_test.txt")
	lines := fileLineScanner("input-data/day6_input.txt")

	world := NewWorld(lines)
	//world.printWorld()
	for i := 0; world.Step(); i++ {
		fmt.Println("Step: ", i+1)
		//world.printWorld()
	}
	fmt.Printf("INFO: guard visited %d distinct locations\n", len(world.visited))

	// while checkWorldExit == false and not in a cycle
	// check next position for cycle (part 2)
	// check next position for collision
	// if collision world.turnRight()
	// check next position for exit
	// if exit
	// world.Step

}
