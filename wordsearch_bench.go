package main

import "testing"

func Benchmarkuper(word string, b *testing.B){
		upercasing(word);
	
}
func Benchmarkuper1(b *testing.B){Benchmarkuper("hello",b)}
func Benchmarkuper2(b *testing.B){Benchmarkuper("the",b)}
func Benchmarkuper3(b *testing.B){Benchmarkuper("what",b)}
func Benchmarkuper4(b *testing.B){Benchmarkuper("is",b)}
func Benchmarkuper5(b *testing.B){Benchmarkuper("do",b)}
func Benchmarkuper6(b *testing.B){Benchmarkuper("and",b)}
func Benchmarkuper7(b *testing.B){Benchmarkuper("there",b)}

