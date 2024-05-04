import { FC, useState } from 'react';

import Layout from '../layouts/Layout';
import useSWR from 'swr';
import { useParams } from 'react-router-dom';
import { useEditor, EditorContent } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'

const formatDateTime = (datetime: string) => {
  return new Date(datetime).toLocaleString();
}

const extensions = [
  StarterKit,
]

const Component: FC = () => {
  const [content, setContent] = useState('')
  const editor = useEditor({
    extensions,
    editorProps: {
      attributes: {
        class: "prose dark:prose-invert prose-sm sm:prose-base lg:prose-lg xl:prose-2xl m-5 focus:outline-none",
      },
    },
  })
  const { short_hash } = useParams<{ short_hash: string }>();

  const { data, error } = useSWR(`/memos/${short_hash}`);

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading...</div>;

  // console.log(data);
  console.log(content);

  return (
    <Layout>
      <div className="py-4 px-2 max-w-7xl mx-auto">
        <div className="grid grid-cols-12 gap-4">
          <div className="col-span-12 lg:col-span-9">
            <div className="bg-[#D8EEFE] shadow border-t-4 border-[#3DA9FC] px-6 md:px-16 py-2 relative">
              <EditorContent editor={editor} value={content} onChange={(e) => setContent(e.target.value)} />
            </div>
          </div>
          <div className="col-span-12 lg:col-span-3">
            <div className="bg-[#D8EEFE] p-2 shadow border-t-4 border-[#3DA9FC] h-48">
              <p>Created At: {formatDateTime(data.created_at)}</p>
              <p>Updated At: {formatDateTime(data.updated_at)}</p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Component;
