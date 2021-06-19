from ctypes import *
import ctypes

lib = cdll.LoadLibrary("./export.so")
lib.my_print()