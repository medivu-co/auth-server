import { useForm } from '@tanstack/react-form';
import type * as React from 'react';
import { Link } from 'react-router';
import * as z from 'zod';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldSeparator } from '@/components/ui/field';
import { Input } from '@/components/ui/input';
import { cn } from '@/lib/utils';

interface LoginFormProps extends React.ComponentProps<'div'> {
  codeChallenge: string;
  codeChallengeMethod: string;
  clientID: string;
  redirectURI: string;
  responseType: string;
  state: string;
}

const loginFormSchema = z.object({
  email: z.string().email('Please enter a valid email address.'),
  password: z.string().min(1, 'Password is required.'),
});

export function LoginForm({ className, ...props }: LoginFormProps) {
  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
    },
    onSubmit: () => {
      const formElement = document.getElementById('login-form') as HTMLFormElement | null;
      formElement?.submit();
    },
    validators: {
      onSubmit: loginFormSchema,
    },
  });

  return (
    <div className={cn('flex flex-col gap-6', className)} {...props}>
      <Card className='overflow-hidden p-0'>
        <CardContent className='grid p-0 md:grid-cols-2'>
          <form
            action='/api/authorize'
            className='p-6 md:p-8'
            id='login-form'
            method='POST'
            onSubmit={(event) => {
              event.preventDefault();
              void form.handleSubmit();
            }}>
            <FieldGroup>
              <div className='flex flex-col items-center gap-2 text-center'>
                <h1 className='text-2xl font-bold'>MEDIVU Account</h1>
                <p className='text-muted-foreground text-balance'>Login to your MEDIVU account</p>
              </div>
              <form.Field name='email'>
                {(field) => {
                  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;
                  return (
                    <Field data-invalid={isInvalid}>
                      <FieldLabel htmlFor={field.name}>Email</FieldLabel>
                      <Input
                        aria-invalid={isInvalid}
                        autoComplete='email'
                        id={field.name}
                        name={field.name}
                        onBlur={field.handleBlur}
                        onChange={(event) => field.handleChange(event.target.value)}
                        placeholder='m@example.com'
                        type='email'
                        value={field.state.value}
                      />
                      {isInvalid && <FieldError errors={field.state.meta.errors} />}
                    </Field>
                  );
                }}
              </form.Field>
              <form.Field name='password'>
                {(field) => {
                  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;
                  return (
                    <Field data-invalid={isInvalid}>
                      <div className='flex items-center'>
                        <FieldLabel htmlFor={field.name}>Password</FieldLabel>
                        <Link
                          className='ml-auto text-sm underline-offset-2 hover:underline'
                          to={`/reset-password${window.location.search}`}>
                          Forgot your password?
                        </Link>
                      </div>
                      <Input
                        aria-invalid={isInvalid}
                        autoComplete='current-password'
                        id={field.name}
                        name={field.name}
                        onBlur={field.handleBlur}
                        onChange={(event) => field.handleChange(event.target.value)}
                        type='password'
                        value={field.state.value}
                      />
                      {isInvalid && <FieldError errors={field.state.meta.errors} />}
                    </Field>
                  );
                }}
              </form.Field>
              <Field>
                <Button type='submit'>Login</Button>
              </Field>
              <input name='response_type' type='hidden' value={props.responseType} />
              <input name='client_id' type='hidden' value={props.clientID} />
              <input name='redirect_uri' type='hidden' value={props.redirectURI} />
              <input name='state' type='hidden' value={props.state} />
              <input name='scope' type='hidden' value='userinfo:write' />
              <input name='code_challenge' type='hidden' value={props.codeChallenge} />
              <input name='code_challenge_method' type='hidden' value={props.codeChallengeMethod} />
              <FieldSeparator className='*:data-[slot=field-separator-content]:bg-card' />
              <FieldDescription className='text-center'>
                Don&apos;t have an account? <Link to={`/signup${window.location.search}`}>Sign up</Link>
              </FieldDescription>
            </FieldGroup>
          </form>
          <div className='bg-muted relative hidden md:block'>
            <img
              alt=''
              className='absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale'
              src='https://ui.shadcn.com/placeholder.svg'
            />
          </div>
        </CardContent>
      </Card>
      <FieldDescription className='px-6 text-center'>
        By clicking login, you agree to our <Link to='/terms'>Terms of Service</Link> and{' '}
        <Link to='/privacy'>Privacy Policy</Link>.
      </FieldDescription>
    </div>
  );
}
