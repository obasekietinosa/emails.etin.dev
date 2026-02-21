import { useEffect, useState } from 'react';
import { api } from '../lib/api';
import { Plus, Key, Copy, Check } from 'lucide-react';
import { format } from 'date-fns';

interface ApiKey {
  id: number;
  name: string;
  key: string;
  created_at: string;
}

export default function ApiKeys() {
  const [keys, setKeys] = useState<ApiKey[]>([]);
  const [loading, setLoading] = useState(true);
  const [newKeyName, setNewKeyName] = useState('');
  const [creating, setCreating] = useState(false);
  const [justCreatedKey, setJustCreatedKey] = useState<string | null>(null);

  useEffect(() => {
    fetchKeys();
  }, []);

  const fetchKeys = async () => {
    try {
      const response = await api.get('/api-keys');
      setKeys(response.data);
    } catch (error) {
      console.error('Failed to fetch API keys', error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    setCreating(true);
    setJustCreatedKey(null);
    try {
      const response = await api.post('/api-keys', { name: newKeyName });
      setKeys([...keys, response.data]);
      setJustCreatedKey(response.data.key);
      setNewKeyName('');
    } catch (error) {
      console.error('Failed to create API key', error);
    } finally {
      setCreating(false);
    }
  };

  const copyKey = (key: string) => {
    navigator.clipboard.writeText(key);
    alert('API Key copied to clipboard!');
  };

  return (
    <div className="space-y-6">
      <div className="md:flex md:items-center md:justify-between">
        <div className="min-w-0 flex-1">
          <h2 className="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
            API Keys
          </h2>
          <p className="mt-1 text-sm text-gray-500">
            Manage API keys for headless mode access.
          </p>
        </div>
      </div>

      <div className="bg-white shadow sm:rounded-lg">
        <div className="px-4 py-5 sm:p-6">
          <h3 className="text-base font-semibold leading-6 text-gray-900">Create new API Key</h3>
          <div className="mt-2 max-w-xl text-sm text-gray-500">
            <p>Enter a name for the new API key.</p>
          </div>
          <form className="mt-5 sm:flex sm:items-center" onSubmit={handleCreate}>
            <div className="w-full sm:max-w-xs">
              <label htmlFor="name" className="sr-only">Name</label>
              <input
                type="text"
                name="name"
                id="name"
                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 pl-2"
                placeholder="e.g. Production Server"
                value={newKeyName}
                onChange={(e) => setNewKeyName(e.target.value)}
                required
              />
            </div>
            <button
              type="submit"
              disabled={creating}
              className="mt-3 inline-flex w-full items-center justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 sm:ml-3 sm:mt-0 sm:w-auto disabled:opacity-50"
            >
              <Plus className="-ml-0.5 mr-1.5 h-5 w-5" aria-hidden="true" />
              Generate
            </button>
          </form>
          {justCreatedKey && (
            <div className="mt-4 rounded-md bg-green-50 p-4">
              <div className="flex">
                <div className="flex-shrink-0">
                  <Check className="h-5 w-5 text-green-400" aria-hidden="true" />
                </div>
                <div className="ml-3">
                  <h3 className="text-sm font-medium text-green-800">API Key Created</h3>
                  <div className="mt-2 text-sm text-green-700">
                    <p>Make sure to copy your API key now.</p>
                    <div className="mt-2 flex items-center space-x-2">
                        <code className="rounded bg-green-100 px-2 py-1 text-green-800 break-all">{justCreatedKey}</code>
                        <button
                            onClick={() => copyKey(justCreatedKey)}
                            className="text-green-600 hover:text-green-800"
                            title="Copy"
                        >
                            <Copy className="h-4 w-4" />
                        </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>

      <div className="bg-white shadow sm:rounded-lg overflow-hidden">
        <ul role="list" className="divide-y divide-gray-200">
          {loading ? (
             <li className="px-4 py-4 sm:px-6 text-center text-sm text-gray-500">Loading...</li>
          ) : keys.length === 0 ? (
             <li className="px-4 py-4 sm:px-6 text-center text-sm text-gray-500">No API keys found.</li>
          ) : (
            keys.map((key) => (
            <li key={key.id} className="px-4 py-4 sm:px-6 hover:bg-gray-50">
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                    <div className="flex-shrink-0">
                        <Key className="h-6 w-6 text-gray-400" aria-hidden="true" />
                    </div>
                    <div className="ml-4">
                        <p className="truncate text-sm font-medium text-indigo-600">{key.name}</p>
                        <p className="mt-1 text-xs text-gray-500">Created on {format(new Date(key.created_at), 'MMM d, yyyy')}</p>
                    </div>
                </div>
                <div className="ml-2 flex flex-shrink-0">
                    <span className="inline-flex rounded-full bg-green-100 px-2 text-xs font-semibold leading-5 text-green-800">
                        Active
                    </span>
                </div>
              </div>
              <div className="mt-2 sm:flex sm:justify-between">
                <div className="sm:flex">
                  <p className="flex items-center text-sm text-gray-500">
                    Key: <code className="ml-1 bg-gray-100 rounded px-1">{key.key}</code>
                  </p>
                </div>
              </div>
            </li>
          ))
          )}
        </ul>
      </div>
    </div>
  );
}
