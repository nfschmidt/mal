
REM ENV_NEW(EO%) -> R%
ENV_NEW:
  REM allocate the data hashmap
  GOSUB HASHMAP

  REM set the data and outer pointer
  ZE%(ZK%)=R%
  ZO%(ZK%)=EO%

  REM update environment pointer and return new environment
  R%=ZK%
  ZK%=ZK%+1
  RETURN

REM ENV_SET(E%, K%, V%) -> R%
ENV_SET:
  HM%=ZE%(E%)
  GOSUB ASSOC1
  ZE%(E%)=R%
  R%=V%
  RETURN

REM ENV_SET_S(E%, K$, V%) -> R%
ENV_SET_S:
  HM%=ZE%(E%)
  GOSUB ASSOC1_S
  ZE%(E%)=R%
  R%=V%
  RETURN

REM ENV_FIND(E%, K%) -> R%
ENV_FIND:
  EF%=E%
  ENV_FIND_LOOP:
    HM%=ZE%(EF%)
    REM More efficient to use GET for value (R%) and contains? (T3%)
    GOSUB HASHMAP_GET
    REM if we found it, save value in T4% for ENV_GET
    IF T3%=1 THEN T4%=R%: GOTO ENV_FIND_DONE
    EF%=ZO%(EF%): REM get outer environment
    IF EF%<>-1 THEN GOTO ENV_FIND_LOOP
  ENV_FIND_DONE:
    R%=EF%
    RETURN

REM ENV_GET(E%, K%) -> R%
ENV_GET:
  GOSUB ENV_FIND
  IF R%=-1 THEN ER%=1: ER$="'" + ZS$(Z%(K%,1)) + "' not found": RETURN
  R%=T4%
  RETURN