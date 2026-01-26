# telegram-bot-api-types

Automatically generate type definition files for the Telegram Bot API in multiple programming languages.

![](https://github.com/TBXark/telegram-bot-api-types/raw/master/preview.jpg)

### Supported Languages

- [x] [TypeScript](./dist/dts)
- [x] [JavaScript](./dist/jsdoc)
- [x] [Json](./dist/spec)
- [x] [Swift](./dist/swift)


### Installation

```sh
go install github.com/tbxark/telegram-bot-api-types@latest
```

### Usage 

You can directly use the precompiled version we provide, or generate your own language version.

```
telegram-bot-api-types --help
  -dist string
        The output directory (default "./dist")
  -help
        Show help
  -lang string
        The output language (default "typescript,jsdoc,spec,swift")
```

## Reference

- Refactored the scraper inspired by [telegram-bot-api-spec](https://github.com/PaulSonOfLars/telegram-bot-api-spec).

## License

**telegram-bot-api-types** is released under the MIT license. [See LICENSE](LICENSE) for details.