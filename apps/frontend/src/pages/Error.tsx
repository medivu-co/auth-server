import { useLocation } from 'react-router-dom';

export function ErrorPage() {
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const snakeCaseError = queryParams.get('error') || 'page_not_found';
  const ErrorMessage = queryParams.get('error_description') || 'Page not found.';

  return (
    <div className='flex min-h-screen items-center justify-center'>
      <div className='text-center'>
        <h1 className='text-6xl font-bold text-gray-800'>
          Error - {toReadableError(snakeCaseError)}
        </h1>
        <p className='mt-4 text-xl text-gray-600'>{ErrorMessage}</p>
        <a
          className='mt-6 inline-block text-indigo-600 hover:text-indigo-500'
          href='/'>
          Go to Home
        </a>
      </div>
    </div>
  );
}

const toReadableError = (snakeCase: string) => {
  return snakeCase
    .split('_')
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ');
};
