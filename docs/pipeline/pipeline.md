# Instruction for Pipeline

## The definition for pipeline

```yaml
name: "The test example"
code: "example_01"
type: "CICD" 
globalVariables:
- key: "key1"
  type: "STRING"
  required: false
  defaultValue: "value1"
```

1. `code` the global unique identifier in all pipeline definition.
2. `name` the display name of pipeline and instance which generate later.
3. `type` The currently supported type is only "CICD".

