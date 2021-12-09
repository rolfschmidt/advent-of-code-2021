package main

import (
    "./helper"
)

func Day2Part1() int64 {
    return Day2Run(false)
}

func Day2Part2() int64 {
    return Day2Run(true)
}

func Day2Run(Part2 bool) int64 {
    distance := int64(0)
    depth := int64(0)
    aim := int64(0)
    for _, line := range helper.ReadFile("day2.txt") {
        vals := helper.Split(line, " ")
        valType := vals[0]
        valVal := helper.String2Int64(vals[1])

        if valType == "forward" {
            distance += valVal
            if Part2 {
                depth += valVal * aim
            }
        } else if valType == "down" {
            if !Part2 {
                depth += valVal
            } else {
                aim += valVal
            }
        } else if valType == "up" {
            if !Part2 {
                depth -= valVal
            } else {
                aim -= valVal
            }
        }
    }

    return distance * depth
}