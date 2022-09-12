using CommandLine.Text;
using CommandLine;

namespace Glow.Commands
{
    [Verb("imagedata", HelpText = "Transform images into base 64 data uris and vice-versa.")]
    class ImageDataOptions
    {
        [Value(0, Required = true, HelpText = "'Convert' to DataUrl or 'Revert' to image files.")]
        public string ?Action { get; set; }

        [Value(1, Required = false, Default=".", HelpText = "Location of the file/directory.")]
        public string ?Path { get; set; }

        [Option('d', "dir", Default = false, HelpText = "Signifies that the location is a directory of images.")]
        public bool Dir { get; set; }

        [Option('m', "max", Default = 100, HelpText = "Max number of images to fit in a single file.")]
        public int Max { get; set; }
    }
}
