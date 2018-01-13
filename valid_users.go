package main

var (
    ValidUsers = make(map[string]int)
)

func init() {
    ValidUsers["o-A1l0zVAgf51kqtZY-oyFGUBi6Y"] = 1
    ValidUsers["o-A1l03mCLRjTP09Z6UZdOVLUBLs"] = 1
)

func IsValidUser(uid string) bool {
    if ok := ValidUsers[uid]; ok {
        return true
    }
    return false
}
