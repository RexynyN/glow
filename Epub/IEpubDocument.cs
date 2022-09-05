using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Glow.Epub
{
    interface IEpubDocument
    {
        public void OutputFile(string path);
    }
}
