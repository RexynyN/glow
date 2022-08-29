import argparse
import os 

# sub-command functions
def foo(args):
    print(args)

def bar(args):
    print('((%s))' % args.z)

# create the top-level parser
parser = argparse.ArgumentParser(description="A multiplatform toolbox for everything you need", prog="wings")
subparsers = parser.add_subparsers()

# Create the parser for the "filerename" command
filerename_parser = subparsers.add_parser('filerename')
filerename_parser.add_argument('dir', type=str, default=os.getcwd())
# filerename_parser.add_argument('y', type=float)
filerename_parser.set_defaults(func=foo)

# create the parser for the "bar" command
parser_bar = subparsers.add_parser('bar')
parser_bar.add_argument('z')
parser_bar.set_defaults(func=bar)

# parse the args and call whatever function was selected
args = parser.parse_args()
args.func(args)