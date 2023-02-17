package main

type inode interface {
	clone() inode
	represent(string)
}
