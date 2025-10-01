export default function NotFoundPage() {
  return (
    <div className='flex min-h-screen items-center justify-center bg-gray-100 dark:bg-gray-900'>
      <div className='text-center'>
        <h1 className='text-6xl font-bold text-gray-800 dark:text-gray-200'>404</h1>
        <p className='mt-4 text-xl text-gray-600 dark:text-gray-400'>Page Not Found!</p>
        <a
          href='/'
          className='mt-6 inline-block text-indigo-600 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300'
        >
          Go to Home
        </a>
      </div>
    </div>
  );
}
