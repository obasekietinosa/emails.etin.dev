import { useEffect, useState } from 'react';
import { api } from '../lib/api';
import { useNavigate, useParams } from 'react-router-dom';
import { ArrowLeft, Save, Loader2 } from 'lucide-react';

interface Template {
  name: string;
  subject: string;
  body: string;
}

export default function TemplateForm() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [template, setTemplate] = useState<Template>({ name: '', subject: '', body: '' });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (id) {
      fetchTemplate(id);
    }
  }, [id]);

  const fetchTemplate = async (templateId: string) => {
    try {
      const response = await api.get(`/templates/${templateId}`);
      setTemplate(response.data);
    } catch (error) {
      console.error('Failed to fetch template', error);
      setError('Failed to load template');
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      if (id) {
        await api.put(`/templates/${id}`, template);
      } else {
        await api.post('/templates', template);
      }
      navigate('/templates');
    } catch (error: any) {
      console.error('Failed to save template', error);
      setError(error.response?.data?.error || 'Failed to save template');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-6">
        {/* Header with back button */}
        <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
                <button
                    onClick={() => navigate('/templates')}
                    className="p-2 text-gray-400 hover:text-gray-500 rounded-full hover:bg-gray-100"
                >
                    <ArrowLeft className="h-6 w-6" />
                </button>
                <h2 className="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
                    {id ? 'Edit Template' : 'New Template'}
                </h2>
            </div>
            <button
                type="submit"
                form="template-form"
                disabled={loading}
                className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 disabled:opacity-50"
            >
                {loading ? <Loader2 className="animate-spin -ml-0.5 mr-1.5 h-5 w-5" /> : <Save className="-ml-0.5 mr-1.5 h-5 w-5" />}
                Save
            </button>
        </div>

        {error && (
            <div className="rounded-md bg-red-50 p-4">
                <div className="flex">
                    <div className="ml-3">
                        <h3 className="text-sm font-medium text-red-800">{error}</h3>
                    </div>
                </div>
            </div>
        )}

        <div className="bg-white shadow sm:rounded-lg">
            <div className="px-4 py-5 sm:p-6">
                <form id="template-form" onSubmit={handleSubmit} className="space-y-6">
                    <div className="grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                        <div className="sm:col-span-4">
                            <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">Template Name</label>
                            <div className="mt-2">
                                <input
                                    type="text"
                                    name="name"
                                    id="name"
                                    required
                                    value={template.name}
                                    onChange={(e) => setTemplate({...template, name: e.target.value})}
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 pl-2"
                                />
                            </div>
                        </div>

                        <div className="sm:col-span-4">
                            <label htmlFor="subject" className="block text-sm font-medium leading-6 text-gray-900">Email Subject</label>
                            <div className="mt-2">
                                <input
                                    type="text"
                                    name="subject"
                                    id="subject"
                                    required
                                    value={template.subject}
                                    onChange={(e) => setTemplate({...template, subject: e.target.value})}
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 pl-2"
                                />
                            </div>
                        </div>

                        <div className="col-span-full">
                            <label htmlFor="body" className="block text-sm font-medium leading-6 text-gray-900">Email Body (HTML supported)</label>
                            <div className="mt-2">
                                <textarea
                                    id="body"
                                    name="body"
                                    rows={10}
                                    required
                                    value={template.body}
                                    onChange={(e) => setTemplate({...template, body: e.target.value})}
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 pl-2 font-mono"
                                />
                            </div>
                            <p className="mt-3 text-sm leading-6 text-gray-600">Use <code>{`{{.VariableName}}`}</code> for variables.</p>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
  );
}
