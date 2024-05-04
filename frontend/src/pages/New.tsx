import { FC } from 'react';

import Layout from '../layouts/Layout';
import useSWR from 'swr';
import { useParams } from 'react-router-dom';

const formatDateTime = (datetime: string) => {
  return new Date(datetime).toLocaleString();
}

const Component: FC = () => {
  return (
    <Layout>
      <div className="py-4 px-2 max-w-7xl mx-auto">
        <div className="grid grid-cols-12 gap-4">
          <div className="col-span-12 lg:col-span-9">
            <div className="bg-[#D8EEFE] p-2 shadow border-t-4 border-[#3DA9FC] p-6 md:px-16">
              <h2 className='text-4xl'>title</h2>
              <p className='py-16'>content</p>
            </div>
          </div>
          <div className="col-span-12 lg:col-span-3">
            <div className="bg-[#D8EEFE] p-2 shadow border-t-4 border-[#3DA9FC] h-48">
              {/* <p>Created At: {formatDateTime(data.created_at)}</p>
              <p>Updated At: {formatDateTime(data.updated_at)}</p> */}
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Component;
