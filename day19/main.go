package main

import (
    "fmt"
    "math"
    "github.com/rolfschmidt/advent-of-code-2021/helper"
)

var scanners []Scanner
var owner int

func main() {
    fmt.Println("Part 1", Part1())
    fmt.Println("Part 2", Part2())
}

func Part1() int {
    return Run(false)
}

func Part2() int {
    return Run(true)
}

type Coord struct {
    x int
    y int
    z int
}

func ContainsCoord(list []Coord, coord Coord) bool {
    for _, check := range list {
        if check.String() == coord.String() {
            return true
        }
    }

    return false
}

func ContainsCoordList(listA []Coord, listB []Coord) bool {
    for _, check := range listB {
        if !ContainsCoord(listA, check) {
            return false
        }
    }

    return true
}

func AppendCoord(list []Coord, coord Coord) []Coord {
    if ContainsCoord(list, coord) {
        return list
    }

    list = append(list, coord)

    return list
}

func RotateList(list []Coord) [][]Coord {
    result := [][]Coord{}
    for ic, coord := range list {
        for ir, rotatedCoord := range coord.Rotations() {
            if ic == 0 {
                result = append(result, []Coord{})
            }
            result[ir] = append(result[ir], rotatedCoord)
        }
    }

    return result
}

func String2Coord(value string) Coord {
    parts := helper.StringArrayInt(helper.Split(value, ","))
    return Coord{
        x: parts[0],
        y: parts[1],
        z: parts[2],
    }
}


func (co *Coord) Add(co2 Coord) Coord {
    return Coord{
        x: co.x + co2.x,
        y: co.y + co2.y,
        z: co.z + co2.z,
    }
}

func (co *Coord) Sub(co2 Coord) Coord {
    return Coord{
        x: co.x - co2.x,
        y: co.y - co2.y,
        z: co.z - co2.z,
    }
}

func (co *Coord) Rotations() []Coord {
    result := []Coord{}

    cX := co.x
    cY := co.y
    cZ := co.z

    result = append(result, Coord{ x: co.x, y: co.y, z: co.z})

    // positive x
    result = append(result, Coord{ x: cX * 1, y: cY * 1, z: cZ * 1})
    result = append(result, Coord{ x: cX * 1, y: cZ * -1, z: cY * 1})
    result = append(result, Coord{ x: cX * 1, y: cY * -1, z: cZ * -1})
    result = append(result, Coord{ x: cX * 1, y: cZ * 1, z: cY * -1})
    // negative x
    result = append(result, Coord{ x: cX * -1, y: cY * -1, z: cZ * 1})
    result = append(result, Coord{ x: cX * -1, y: cZ * 1, z: cY * 1})
    result = append(result, Coord{ x: cX * -1, y: cY * 1, z: cZ * -1})
    result = append(result, Coord{ x: cX * -1, y: cZ * -1, z: cY * -1})
    // positive y
    result = append(result, Coord{ x: cY * 1, y: cZ * 1, z: cX * 1})
    result = append(result, Coord{ x: cY * 1, y: cX * -1, z: cZ * 1})
    result = append(result, Coord{ x: cY * 1, y: cZ * -1, z: cX * -1})
    result = append(result, Coord{ x: cY * 1, y: cX * 1, z: cZ * -1})
    // negative y
    result = append(result, Coord{ x: cY * -1, y: cZ * -1, z: cX * 1})
    result = append(result, Coord{ x: cY * -1, y: cX * 1, z: cZ * 1})
    result = append(result, Coord{ x: cY * -1, y: cZ * 1, z: cX * -1})
    result = append(result, Coord{ x: cY * -1, y: cX * -1, z: cZ * -1})
    // positive z
    result = append(result, Coord{ x: cZ * 1, y: cX * 1, z: cY * 1})
    result = append(result, Coord{ x: cZ * 1, y: cY * -1, z: cX * 1})
    result = append(result, Coord{ x: cZ * 1, y: cX * -1, z: cY * -1})
    result = append(result, Coord{ x: cZ * 1, y: cY * 1, z: cX * -1})
    // negative z
    result = append(result, Coord{ x: cZ * -1, y: cX * -1, z: cY * 1})
    result = append(result, Coord{ x: cZ * -1, y: cY * 1, z: cX * 1})
    result = append(result, Coord{ x: cZ * -1, y: cX * 1, z: cY * -1})
    result = append(result, Coord{ x: cZ * -1, y: cY * -1, z: cX * -1})

    return result
}

func (co *Coord) Directions() []Coord {
    result := []Coord{}
    for _, dX := range []int{1, -1} {
        for _, dY := range []int{1, -1} {
            for _, dZ := range []int{1, -1} {
                check := Coord{
                    x: co.x * dX,
                    y: co.y * dY,
                    z: co.z * dZ,
                }
                result = append(result, check)
            }
        }
    }
    return result
}

func (co *Coord) Variants(co2 Coord) []Coord {
    result := []Coord{}

    bAdd := co2.Add(*co)
    bSub := co2.Sub(*co)
    for _, wantFindX := range []int{ bAdd.x, bSub.x } {
        for _, wantFindY := range []int{ bAdd.y, bSub.y } {
            for _, wantFindZ := range []int{ bAdd.z, bSub.z } {
                check := Coord{
                    x: wantFindX,
                    y: wantFindY,
                    z: wantFindZ,
                }
                result = append(result, check)
            }
        }
    }

    return result
}

func (co Coord) Sum() int {
    return co.x + co.y + co.z
}

func (co Coord) Distance(co2 Coord) Coord {
    return Coord{
        x: int(math.Abs(float64(helper.IntMax(co.x, co2.x) - helper.IntMin(co.x, co2.x)))),
        y: int(math.Abs(float64(helper.IntMax(co.y, co2.y) - helper.IntMin(co.y, co2.y)))),
        z: int(math.Abs(float64(helper.IntMax(co.z, co2.z) - helper.IntMin(co.z, co2.z)))),
    }
}

func (co Coord) String() string {
    return helper.Int2String(co.x) + "," + helper.Int2String(co.y) + "," + helper.Int2String(co.z)
}

type Scanner struct {
    name string
    pos Coord
    known bool
    beacons []Coord
    beaconsRelativeAbsCount map[string][]Coord
}

func ListRelatives(list []Coord) map[string][]Coord {
    result := map[string][]Coord{}
    for _, cA := range list {
        for _, cB := range list {
            if cA.String() == cB.String() {
                continue
            }

            key := cA.Distance(cB).String()
            result[key] = append(result[key], cA)
        }
    }

    return result
}

func (sc *Scanner) OverlapCoordsMap(sc2 Scanner, work func(bA string, bACoords []Coord, bBCoords []Coord) bool ) {
    for bA, bACoords := range sc.beaconsRelativeAbsCount {
        bBCoords, ok := sc2.beaconsRelativeAbsCount[bA]
        if !ok {
            continue
        }

        if len(bACoords) != len(bBCoords) {
            continue
        }

        continueFunc := work(bA, bACoords, bBCoords)
        if !continueFunc {
            break
        }
    }
}

func (sc *Scanner) Overlap(sc2 Scanner) ([]Coord, []Coord) {
    overlapCountRequired := 12
    overlapCountRequired = (overlapCountRequired - 1) * overlapCountRequired

    relativesListA := ListRelatives(sc.beacons)

    hotRelativesA := []Coord{}
    hotRelativesB := []Coord{}
    for _, list := range RotateList(sc2.beacons) {

        matchRelativesA := []Coord{}
        matchRelativesB := []Coord{}
        for relativesKey, relativesB := range ListRelatives(list) {
            relativesA, ok := relativesListA[relativesKey]
            if !ok {
                continue
            }

            if len(relativesA) != len(relativesB) {
                continue
            }

            for _, coord := range relativesA {
                matchRelativesA = AppendCoord(matchRelativesA, coord)
            }
            for _, coord := range relativesB {
                matchRelativesB = AppendCoord(matchRelativesB, coord)
            }
        }

        if len(matchRelativesB) > len(hotRelativesB) {
            hotRelativesA = matchRelativesA
            hotRelativesB = matchRelativesB
        }
    }

    if len(hotRelativesB) >= 12 {
        return hotRelativesA, hotRelativesB
    }

    return hotRelativesA, hotRelativesB
}

func (sc *Scanner) SetPos(sc2 *Scanner, hotRelativesA []Coord, hotRelativesB []Coord) {

    counts := map[string]int{}
    for _, coordB := range hotRelativesB {
        for _, coordA := range hotRelativesA {
            for _, variant := range coordB.Variants(coordA) {
                counts[variant.String()] += 1
            }
        }
    }

    highestPos := ""
    highestCount := 0
    for key, count := range counts {
        if count > highestCount {
            highestPos = key
            highestCount = count
        }
    }

    highestCoord := String2Coord(highestPos)
    sc2.pos = highestCoord

    newCoords := map[string][]Coord{}
    for _, coord := range sc2.beacons {
        for ri, roCoord := range coord.Rotations() {
            for hi, hCoord := range highestCoord.Directions() {
                newAdd := roCoord.Add(hCoord)
                key := helper.Int2String(ri) + "_" + helper.Int2String(hi)
                if _, ok := newCoords[key]; !ok {
                    newCoords[key] = []Coord{}
                }

                newCoords[key] = append(newCoords[key], newAdd)
            }
        }
    }

    FIND:
    for _, list := range newCoords {
        if ContainsCoordList(list, hotRelativesA) {
            sc2.beacons = list
            break FIND
        }
    }

    for _, coord := range sc.beacons {
        sc2.beacons = AppendCoord(sc2.beacons, coord)
    }
}

func setupScanner() {
    content := helper.ReadFileString("input.txt")
    contentScanners := helper.Split(content, "\n\n")

    for _, scanner := range contentScanners {
        beacons := helper.Split(scanner, "\n")
        name := helper.Trim(helper.Replace(beacons[0], "---", ""))
        beacons = beacons[1:]

        scanner := Scanner{ name: name }
        for _, beaconstr := range beacons {
            beacons := helper.StringArrayInt(helper.Split(beaconstr, ","))
            scanner.beacons = append(scanner.beacons, Coord{ x: beacons[0], y: beacons[1], z: beacons[2], })
        }

        scanners = append(scanners, scanner)
    }

    done := map[int]bool{}
    for len(done) < len(scanners) {
        SCANNER:
        for s1 := range scanners {
            sc1 := &scanners[s1]
            if s1 != owner {
                continue
            }

            for s2 := range scanners {
                sc2 := &scanners[s2]
                if sc1.name == sc2.name || done[s2] {
                    continue
                }

                hotRelativesA, hotRelativesB := sc1.Overlap(*sc2)
                if len(hotRelativesB) >= 12 {
                    // fmt.Println(sc1.name, "overlaps with", sc2.name, " with beacon count", len(sc1.beacons))
                    sc1.SetPos(sc2, hotRelativesA, hotRelativesB)
                    owner = s2
                    done[s1] = true
                    done[s2] = true

                    // break LOOP
                    continue SCANNER
                }
            }
        }
    }
}

func Run(Part2 bool) int {
    if len(scanners) < 1 {
        setupScanner()
    }

    if Part2 {
        hotDistance := -1
        for _, sA := range scanners {
            for _, sB := range scanners {
                if sA.name == sB.name {
                    continue
                }

                distance := sA.pos.Distance(sB.pos).Sum()
                hotDistance = helper.IntMax(hotDistance, distance)
            }
        }

        return hotDistance
    }

    return len(scanners[owner].beacons)
}
