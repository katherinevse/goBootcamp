package utils

import "github.com/gorilla/mux"

//ресурс это что такое?

utils.BuildbookResoursePrefix(router, bookResoursePrefix)
utils.BuildmanybookResoursePrefix(router, manyBooksResoursePrefix)


func BuildbookResoursePrefix(router *mux.Router, prefix string)