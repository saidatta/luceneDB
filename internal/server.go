package main

func newServer(hostname string, port string) (*database, error) {
	d := database{data: nil, port: port}
	var err error
	d.data, err = newDatabase(hostname)
	if err != nil {
		return nil, err
	}

	d.index, err = newDatabase(hostname + ".index")

	return &d, err
}

func reindex() {

}
