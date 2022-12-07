package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

// $ cd /
// $ ls

type Directory struct {
    directories map[string]int
    files map[string]int
    total_size int
    parent int
}

func compute_total_size(directories []Directory, index int) int {
    // compute size of children
    result := 0
    for _, child_index := range directories[index].directories {
        result += compute_total_size(directories, child_index)
        directories[index].total_size += directories[child_index].total_size
    }
    if directories[index].total_size < 100000 {
        result += directories[index].total_size
    }
    return result
}

func find_smallest_directory(directories []Directory, required_space int, index int) int {
    // check if children have a result
    result := -1
    for _, child_index := range directories[index].directories {
        child_result := find_smallest_directory(directories, required_space, child_index)
        if child_result >= 0 {
            if result < 0 {
                result = child_result
            } else if directories[result].total_size > directories[child_result].total_size {
                result = child_result
            }
        }
    }
    if result < 0 {
        // no result yet
        if directories[index].total_size >= required_space {
            result = index
        }
    }
    return result
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    read_ls_output := false

    directories := make([]Directory, 0, 10)
    var root Directory
    root.parent = -1
    root.total_size = 0
    root.directories = map[string]int{}
    root.files = map[string]int{}
    directories = append(directories, root)
    current_node := 0
    for scanner.Scan() {
        line := scanner.Text()
        if line[0] == '$' {
            read_ls_output = false
            // command
            s := strings.Split(line, " ")
            if s[1] == "cd" {
                // cd
                if s[2] == ".." {
                    // move up the tree
                    if directories[current_node].parent < 0 {
                        panic("already in root")
                    }
                    current_node = directories[current_node].parent
                } else if s[2] == "/" {
                    // move to root
                    current_node = 0
                } else {
                    new_val, ok := directories[current_node].directories[s[2]]
                    if !ok {
                        fmt.Println(directories[current_node].directories)
                        panic("subdir not found")
                    }
                    current_node = new_val
                }
            } else if s[1] == "ls" {
                // ls
                read_ls_output = true
            } else {
                panic("Unexpected keyword")
            }
        } else if read_ls_output {
            s := strings.Split(line, " ")
            if s[0] == "dir" {
                // directory
                dirname := s[1]
                if _, ok := directories[current_node].directories[dirname]; ok {
                    panic("directory already found")
                }
                var new_dir Directory
                new_dir.parent = current_node
                new_dir.total_size = 0
                new_dir.directories = map[string]int{}
                new_dir.files = map[string]int{}
                new_index := len(directories)
                directories = append(directories, new_dir)
                directories[current_node].directories[dirname] = new_index
            } else {
                size, err := strconv.Atoi(s[0])
                if err != nil {
                    log.Fatal(err)
                }
                filename := s[1]
                if _, ok := directories[current_node].files[filename]; ok {
                    panic("filename already found")
                }
                directories[current_node].files[filename] = size
                directories[current_node].total_size += size
            }
        } else {
            panic("Expected ls output")
        }
    }
    compute_total_size(directories, 0)

    fs_size  := 70000000
    required := 30000000
    current_space := fs_size - directories[0].total_size

    result := find_smallest_directory(directories, required - current_space, 0)
    fmt.Println(directories[result].total_size)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
