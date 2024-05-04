import { Routes, Route } from 'react-router-dom';
import { SWRConfig } from 'swr';
import { fetchInstance } from './utils/fetchInstance';

import Home from './pages/Home';
import Detail from './pages/Detail';
import New from './pages/New';

function App() {
  return (
    <SWRConfig
      value={{
        fetcher: (resource, init) =>
          fetchInstance()
            .get(resource, init)
            .then((res) => res.data),
      }}
    >
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/new" element={<New />} />
        <Route path="/:short_hash" element={<Detail />} />
      </Routes>
    </SWRConfig>
  );
}

export default App;
