import { ChangeEvent, FC, useEffect, useState } from 'react';

import Layout from '../layouts/Layout';
import useSWRInfinite from 'swr/infinite';

const sortOptions = [
  { value: 'id_asc', label: 'ID Asc' },
  { value: 'id_desc', label: 'ID Desc' },
  { value: 'created_at_asc', label: 'Created At Asc' },
  { value: 'created_at_desc', label: 'Created At Desc' },
  { value: 'updated_at_asc', label: 'Updated At Asc' },
  { value: 'updated_at_desc', label: 'Updated At Desc' },
  { value: 'view_asc', label: 'View Asc' },
  { value: 'view_desc', label: 'View Desc' },
];

const sort = (data: any[], option: string) => {
  localStorage.setItem('sortOption', option);
  switch (option) {
    case 'id_asc':
      return data.sort((a, b) => a.id - b.id);
    case 'id_desc':
      return data.sort((a, b) => b.id - a.id);
    case 'created_at_desc':
      return data.sort(
        (a, b) =>
          new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
      );
    case 'created_at_asc':
      return data.sort(
        (a, b) =>
          new Date(a.created_at).getTime() - new Date(b.created_at).getTime(),
      );
    case 'updated_at_asc':
      return data.sort(
        (a, b) =>
          new Date(a.updated_at).getTime() - new Date(b.updated_at).getTime(),
      );
    case 'updated_at_desc':
      return data.sort(
        (a, b) =>
          new Date(b.updated_at).getTime() - new Date(a.updated_at).getTime(),
      );
    case 'view_asc':
      return data.sort((a, b) => a.view_count - b.view_count);
    case 'view_desc':
      return data.sort((a, b) => b.view_count - a.view_count);
    default:
      return data;
  }
};

const pinned = (data: any[]) => {
  const pinned = data.filter((item) => item.pinned);
  const notPinned = data.filter((item) => !item.pinned);
  return [...pinned, ...notPinned];
};

const Component: FC = () => {
  const [sortOption, setSortOption] = useState(
    localStorage.getItem('sortOption') || 'created_at_desc',
  );
  const handleSort = (e: ChangeEvent<HTMLSelectElement>) => {
    setSortOption(e.target.value);
  };

  const { data, error, size, setSize } = useSWRInfinite(
    (index: number) => `/memos?page=${index}&per=50`,
    {
      revalidateIfStale: false,
      revalidateOnFocus: false,
      revalidateFirstPage: false,
    },
  );
  const isEnd = data && data[data.length - 1].length === 0;

  useEffect(() => {
    const handleScroll = () => {
      if (
        window.innerHeight + window.scrollY >= document.body.scrollHeight &&
        !isEnd
      ) {
        setSize(size + 1);
      }
    };
    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  });

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading...</div>;

  return (
    <Layout>
      <div className="py-4 px-2 max-w-7xl mx-auto">
        <div className="grid grid-cols-12 gap-4">
          <div className="col-span-12">
            <div className="flex justify-between items-center">
              <p className="font-bold">All Memos</p>
              <div className="flex items-center space-x-2">
                <p>Sort by:</p>
                <select
                  name="sort"
                  id="sort"
                  onChange={handleSort}
                  className="border border-gray-300 rounded-lg active:outline-none focus:outline-none"
                  value={sortOption}
                >
                  {sortOptions.map((item) => (
                    <option key={item.value} value={item.value}>
                      {item.label}
                    </option>
                  ))}
                </select>
              </div>
            </div>
          </div>
          {pinned(sort(data.flat(), sortOption)).map((item) => (
            <a
              key={item.id}
              href={item.short_hash}
              className="col-span-6 md:col-span-3 lg:col-span-2 relative hover:shadow-lg transition duration-300 ease-in-out cursor-pointer"
            >
              <div className="bg-[#D8EEFE] p-2 shadow border-t-4 border-[#3DA9FC] h-48">
                <p className="font-bold break-all line-clamp-3">{item.title}</p>
                <p className="break-all line-clamp-3">{item.content}</p>
              </div>
              {item.pinned && (
                <div className="absolute bottom-0 right-0 border-8 border-transparent border-b-[#3DA9FC] border-r-[#3DA9FC]" />
              )}
            </a>
          ))}
        </div>
      </div>
    </Layout>
  );
};

export default Component;
