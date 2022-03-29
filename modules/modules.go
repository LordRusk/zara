package modules

/* this is an interface to interface with zara */
type Module interface {
	Run() (output string, err error)
	Int() int /* returns update intervel. */
	Sig() int /* returns update signal */

	/*
	 * simply stores an int for modules position
	 * in the bar. See Time for implementation
	 */
	SetPos(int) /* sets position */
	Pos() int   /* returns position */

	Name() string /* needed for error identification */
}
