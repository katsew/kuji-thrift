 #!/bin/sh

DEFAULT_TARGET=
DEFAULT_LANG=go

TARGET=$DEFAULT_TARGET
LANG=$DEFAULT_LANG

SRC_DIR=./thrift/source
DEST_DIR=./thrift/dist

cd `dirname $0`
cd ../../

if [ ! -z "$1" ]; then
    LANG=$1
fi

if [ ! -z "$2" ]; then
    TARGET=$2
fi

echo "language: $LANG"
echo "target: $TARGET"

if [ -z "$TARGET" ]; then
    echo "No target specified!"
    exit 1
fi

if [ ! -d "$DEST_DIR/$LANG" ]; then
    echo "Directory not exists, create $LANG dir."
    mkdir -p $DEST_DIR/$LANG
fi

if [ ! -d "$DEST_DIR/$LANG/$TARGET" ]; then
    echo "Directory not exists, create $LANG/$TARGET dir."
    mkdir -p $DEST_DIR/$LANG/$TARGET
fi

if [ $LANG == go ]; then

echo "output lang: go"
thrift -I $SRC_DIR/$TARGET -out $DEST_DIR/$LANG -r --gen $LANG  $SRC_DIR/$TARGET/*.thrift

elif [ $LANG == php ]; then

echo "output lang: php"
thrift -I $SRC_DIR/$TARGET -out $DEST_DIR/$LANG -r --gen $LANG $SRC_DIR/$TARGET/*.thrift

elif [ $LANG == ruby ]; then

echo "output lang: ruby"
thrift -I $SRC_DIR/$TARGET -out $DEST_DIR/$LANG/$TARGET -r --gen rb  $SRC_DIR/$TARGET/*.thrift

elif [ $LANG == java ]; then

thrift -I $SRC_DIR/$TARGET -out $DEST_DIR/$LANG -r --gen $LANG  $SRC_DIR/$TARGET/*.thrift

else

echo "Not implemented yet."

fi

echo "Compile task finished!"