using System;
using System.Collections.Generic;
using System.IO.Compression;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Glow.Epub
{
    class Epub
    {
        public EpubContainer Content{ get; set; }
    
        public void CreateEpub(string epubPath)
        {
            string startPath = Path.Combine(Directory.GetCurrentDirectory(), "_epub_");

            ZipFile.CreateFromDirectory(startPath, epubPath);
        }
    }
}
