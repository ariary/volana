package volana

type Historic struct {
	Previous []string //LIFO
	Next     []string //FIFO
	Current  string
}

//Return previous command and synchronize Next commands historical accordingly
func (h *Historic) GetPrevious() (command string) {

	//Update Next
	if h.Current != "" {
		h.Next = append([]string{h.Current}, h.Next...)
	}

	//Update Current
	n := len(h.Previous) - 1
	if n >= 0 {
		h.Current = h.Previous[n]
	} else {
		h.Current = ""
	}

	//Update Previous
	if n >= 0 {
		h.Previous = h.Previous[:n] // ~pop
	}

	//Return current
	command = h.Current
	return command
}

//Return next command and synchronize previous commands historical accordingly
func (h *Historic) GetNext() (command string) {
	//Update Previous
	if h.Current != "" {
		h.Previous = append(h.Previous, h.Current)
	}

	//Update Current
	n := len(h.Next) - 1
	if n >= 0 {
		h.Current = h.Next[0]
	}

	//Update Next
	if n >= 0 {
		h.Next = h.Next[1:]
	}

	//Return current
	command = h.Current
	return command
}

func (h *Historic) Add(command string) {
	h.Previous = append(h.Previous, command)
}
