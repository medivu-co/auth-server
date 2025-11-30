import { useSearchParams } from 'react-router';
import { LoginForm } from '@/components/login-form';

export function AuthorizePage() {
  const [searchParams] = useSearchParams();
  const responseType = searchParams.get('response_type') || 'code';
  const clientID = searchParams.get('client_id');
  const rawRedirectURI = searchParams.get('redirect_uri');
  const state = searchParams.get('state');
  const codeChallenge = searchParams.get('code_challenge');
  const codeChallengeMethod = searchParams.get('code_challenge_method');

  if (!clientID || !rawRedirectURI || !state || !codeChallenge || !codeChallengeMethod) {
    // log which parameters are missing
    const missingParams = [];
    if (!clientID) missingParams.push('client_id');
    if (!rawRedirectURI) missingParams.push('redirect_uri');
    if (!state) missingParams.push('state');
    if (!codeChallenge) missingParams.push('code_challenge');
    if (!codeChallengeMethod) missingParams.push('code_challenge_method');
    console.error(`Missing required query parameters: ${missingParams.join(', ')}`);
    return (
      <div className='flex min-h-screen items-center justify-center bg-gray-100 dark:bg-gray-900'>
        <div className='text-center'>
          <h1 className='text-2xl font-bold text-gray-800 dark:text-gray-200'>Invalid Request!</h1>
          <p className='mt-4 text-gray-600 dark:text-gray-400'>Missing required query parameters.</p>
        </div>
      </div>
    );
  }
  const redirectURI = decodeURIComponent(rawRedirectURI);
  console.log({
    clientID,
    codeChallenge,
    codeChallengeMethod,
    redirectURI,
    responseType,
    state,
  });

  return (
    <div className='bg-muted flex min-h-svh flex-col items-center justify-center p-6 md:p-10'>
      <div className='w-full max-w-sm md:max-w-4xl'>
        <LoginForm
          clientID={clientID}
          codeChallenge={codeChallenge}
          codeChallengeMethod={codeChallengeMethod}
          redirectURI={rawRedirectURI}
          responseType={responseType}
          state={state}
        />
      </div>
    </div>
  );
}
