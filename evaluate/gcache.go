package evaluate


func main() {
	gc := gcache.New(10).
		LFU().
		Build()
	gc.Set("key", "ok")

	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", v)
}