package skills

// LoadBuiltins returns embedded skills compiled into the binary.
// Currently empty — all skills are loaded from disk at runtime
// (user-level ~/.ashecode/skills/ or project-level .ashecode/skills/).
func LoadBuiltins() []*Skill {
	return nil
}
