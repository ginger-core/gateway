package instruction

func (i *instruction) WithField(f Field) Instruction {
	if i.fieldNameMap == nil {
		i.fieldNameMap = make(map[string]Field)
		i.fieldKeyMap = make(map[string]Field)
	}
	i.fieldNameMap[f.GetName()] = f
	i.fieldKeyMap[f.GetKey()] = f
	return i
}

func (i *instruction) GetFields() map[string]Field {
	return i.fieldNameMap
}

func (i *instruction) GetField(name string) Field {
	if i.fieldNameMap == nil {
		return nil
	}
	return i.fieldNameMap[name]
}

func (i *instruction) GetFieldByKey(key string) Field {
	if i.fieldKeyMap == nil {
		return nil
	}
	return i.fieldKeyMap[key]
}
