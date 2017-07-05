# continuous-deployment-templates

Template, modules, etc for doing demos/hacks for CD

---

[Guide](wiki/README.md)

## Project Structure

```
# Terraform Templates. Modules used to make environments
├── infrastructure
│   └── modules

# generate: Source library
├── src
│   └── templates

# Templates for generate terraform files.
├── templates
│   ├── ciServer
│   │   └── templates
│   ├── simpleEmbedded
│   │   └── files

# Test Libraries
├── testing
│   └── helper
├── tests
│   ├── rancher-db-setup
│   └── rancher-quick-setup

# Wiki/Docs
└── wiki
```

### Generate

Generates Terraform configurations using input from a series of prompts.
