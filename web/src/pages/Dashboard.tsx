import { useEffect, useState } from 'react';
import { api } from '../lib/api';
import { format } from 'date-fns';
import { Link } from 'react-router-dom';
import { Plus, ArrowRight, Activity } from 'lucide-react';

interface EmailLog {
  id: number;
  recipient: string;
  subject: string;
  status: string;
  mode: string;
  created_at: string;
}

export default function Dashboard() {
  const [logs, setLogs] = useState<EmailLog[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchLogs = async () => {
      try {
        const response = await api.get('/logs');
        // Take top 5
        setLogs(response.data.slice(0, 5));
      } catch (error) {
        console.error('Failed to fetch logs', error);
      } finally {
        setLoading(false);
      }
    };

    fetchLogs();
  }, []);

  return (
    <div className="space-y-6">
      <div className="md:flex md:items-center md:justify-between">
        <div className="min-w-0 flex-1">
          <h2 className="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
            Dashboard
          </h2>
        </div>
        <div className="mt-4 flex md:ml-4 md:mt-0">
          <Link
            to="/templates"
            className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            <Plus className="-ml-0.5 mr-1.5 h-5 w-5" aria-hidden="true" />
            New Template
          </Link>
        </div>
      </div>

      {/* Stats or Recent Activity */}
      <div className="overflow-hidden rounded-lg bg-white shadow">
        <div className="p-6">
          <h3 className="text-base font-semibold leading-6 text-gray-900 flex items-center">
            <Activity className="h-5 w-5 mr-2 text-gray-500" />
            Recent Activity
          </h3>
          <div className="mt-6 flow-root">
            <ul role="list" className="-my-5 divide-y divide-gray-200">
              {loading ? (
                <li className="py-5 text-center text-sm text-gray-500">Loading activity...</li>
              ) : logs.length === 0 ? (
                <li className="py-5 text-center text-sm text-gray-500">No recent activity found.</li>
              ) : (
                logs.map((log) => (
                  <li key={log.id} className="py-4">
                    <div className="flex items-center space-x-4">
                      <div className="min-w-0 flex-1">
                        <p className="truncate text-sm font-medium text-gray-900">
                          Sent to <span className="font-bold">{log.recipient}</span>
                        </p>
                        <p className="truncate text-sm text-gray-500">
                          Subject: {log.subject}
                        </p>
                      </div>
                      <div className="inline-flex items-center text-sm text-gray-500">
                        <span className={`inline-flex items-center rounded-md px-2 py-1 text-xs font-medium ring-1 ring-inset ${
                            log.mode === 'managed' ? 'bg-blue-50 text-blue-700 ring-blue-600/20' : 'bg-purple-50 text-purple-700 ring-purple-600/20'
                        } mr-2`}>
                            {log.mode}
                        </span>
                        {format(new Date(log.created_at), 'MMM d, yyyy HH:mm')}
                      </div>
                    </div>
                  </li>
                ))
              )}
            </ul>
          </div>
          <div className="mt-6">
            <Link
              to="/logs"
              className="flex w-full items-center justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
            >
              View all
              <ArrowRight className="ml-2 h-4 w-4 text-gray-400" />
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
