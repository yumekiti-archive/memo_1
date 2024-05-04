import { FC, ReactNode, memo, useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';

type Props = {
  children: ReactNode;
};

const Component: FC<Props> = ({ children }) => {
  const [searchParams] = useSearchParams();
  const q = searchParams.get('q') || '';
  const userNavigate = useNavigate();
  const [searchText, setSearchText] = useState<string>(q);

  const handleSearch = () => {
    if (searchText === '') return;
    userNavigate(`/search?q=${searchText}`);
  };

  const handleCreate = () => {
    userNavigate('/new');
  }

  return (
    <div className="App h-screen flex flex-col">
      <header>
        <nav className="bg-[#D8EEFE] py-2 flex justify-around items-center gap-2 md:gap-0 px-2 md:px-0">
          <a href="/" className="hidden md:block hover:underline text-xl">
            HOME
          </a>
          <div className="flex items-center gap-2 w-full md:w-[40%] lg:w-[30%]">
            <input
              type="text"
              name="search"
              id="search"
              placeholder="Search"
              className="px-2 py-2 border border-gray-300 rounded-lg w-full active:outline-none focus:outline-none"
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === 'Enter') handleSearch();
              }}
            />
            <button
              className="p-2 flex items-center justify-center bg-[#3DA9FC] text-white rounded-lg active:outline-none focus:outline-none"
              aria-label="SearchButton"
              onClick={handleSearch}
            >
              Search
            </button>
          </div>
          <button className="bg-[#3DA9FC] text-white px-4 py-2 rounded-lg" onClick={handleCreate}>
            CREATE
          </button>
        </nav>
      </header>
      <main className="flex-grow">{children}</main>
    </div>
  );
};

export default memo(Component);
