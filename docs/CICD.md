# CI/CD

## Continuous integration

### Build

| Permissions  | Triggers       |
| ------------ | -------------- |
| Read / Write | Push to `main` |

```mermaid
stateDiagram-v2
	direction LR
	
	[*] --> Build
	Build --> Test
	Test --> Versioning
	Versioning --> Release
	Release --> [*]
```

### Coverage

| Permissions | Triggers       |
| ----------- | -------------- |
| Read only   | Push to `main` |

```mermaid
stateDiagram-v2
	direction LR
	
	[*] --> Test
	Test --> Coverage
	Coverage --> [*]
	
```
