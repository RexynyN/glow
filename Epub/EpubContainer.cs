using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Glow.Epub
{
    class EpubContainer
    {
        private List<IEpubDocument> _documents;

        public EpubContainer(List<IEpubDocument> documents)
        {
            _documents = documents;
        }

        public void AddDocument(IEpubDocument epubDoc)
        {
            _documents.Add(epubDoc);
        }


    }
}
