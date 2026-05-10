package warehouses

type Dto struct {
	Code     string
	Name     string
	Address  string
	Timezone string
	Active   bool
}

type CreateDto struct {
	Code     string
	Name     string
	Address  string
	Timezone string
	Active   bool
}

type UpdateDto struct {
	Name     *string
	Address  *string
	Timezone *string
	Active   *bool
}
