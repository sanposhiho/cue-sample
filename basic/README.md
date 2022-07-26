# Basic samples

## CUE から YAML を生成

```
$ cue export --out yaml 
User:
  TEL: "123-1234-1234"
  Age: 10
  Is_activated: true
```

## YAML から CUE を生成

```
$ cue import ./testdata/...  -f
```

## YAML を CUE を用いてバリデーション

**バリデーション成功**

```
$ cue vet . ./testdata/valid.yaml
```

**バリデーション失敗**

```
$ cue vet . ./testdata/invalid.yaml
User.Age: conflicting values 10 and 100000000:
    ./data.cue:5:7
    ./testdata/invalid.yaml:2:40
User.Is_activated: conflicting values true and false:
    ./data.cue:6:16
    ./testdata/invalid.yaml:3:36
User.TEL: conflicting values "123-1234-1234" and "123-1234-123":
    ./data.cue:4:7
    ./testdata/invalid.yaml:2:9
User._length: 1 errors in empty disjunction:
```