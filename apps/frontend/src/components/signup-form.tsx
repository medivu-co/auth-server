import { useForm } from '@tanstack/react-form';
import axios from 'axios';
import type * as React from 'react';
import { Link, useNavigate } from 'react-router';
import { toast } from 'sonner';
import * as z from 'zod';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel, FieldSeparator } from '@/components/ui/field';
import { Input } from '@/components/ui/input';
import { cn } from '@/lib/utils';

const signUpFormSchema = z.object({
  name: z
    .string()
    .trim()
    .min(1, 'Please enter your name.')
    .max(25, 'Name must be at most 25 characters.'),
  confirmPassword: z.string().min(1, 'Please confirm your password.'),
  email: z.email('Please enter a valid email address.'),
  password: z
    .string()
    .regex(/[\x20-\x7E]+/, {
      message: 'Password must contain only printable ASCII characters (alphabets, numbers, symbols).',
    })
    .min(12, 'Password must be at least 12 characters.')
    .max(64, 'Password must be at most 64 characters.'),
});

export function SignupForm({ className, ...props }: React.ComponentProps<'div'>) {
  const navigate = useNavigate();
  const form = useForm({
    defaultValues: {
      name: '',
      confirmPassword: '',
      email: '',
      password: '',
    },
    onSubmit: async ({ value }) => {
      console.log('Form submitted with values:', value);
      try {
        const response = await axios.post('/api/signup', value);
        console.log('Sign-up successful:', response.data);
        navigate('/pending-verification', { replace: true });
      } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
          // Handle known errors from the server
          switch (error.response.status) {
            case 400:
              toast.error('Invalid input. Please check your details and try again.');
              break;
            case 409:
              toast.error('Email already exists. Please use a different email.');
              break;
            case 500:
              toast.error('Server error. Please try again later.');
              break;
            default:
              toast.error('An unexpected error occurred. Please try again later.');
          }
        } else {
          // Handle unexpected errors
          toast.error('An unexpected error occurred. Please try again later.');
        }
        console.error('Sign-up failed:', error);
      }
    },
    validators: {
      onSubmit: signUpFormSchema,
    },
  });
  return (
    <div className={cn('flex flex-col gap-6', className)} {...props}>
      <Card className='overflow-hidden p-0'>
        <CardContent className='grid p-0 md:grid-cols-2'>
          <form
            className='p-6 md:p-8'
            id='signup-form'
            onSubmit={(e) => {
              e.preventDefault();
              form.handleSubmit();
            }}>
            <FieldGroup>
              <div className='flex flex-col items-center gap-2 text-center'>
                <h1 className='text-2xl font-bold'>Create your MEDIVU account</h1>
                <p className='text-muted-foreground text-sm text-balance'>
                  Enter your details below to create your account
                </p>
              </div>
              <form.Field name='name'>
                {(field) => {
                  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;
                  return (
                    <Field data-invalid={isInvalid}>
                      <FieldLabel htmlFor={field.name}>Name</FieldLabel>
                      <Input
                        aria-invalid={isInvalid}
                        autoComplete='name'
                        id={field.name}
                        maxLength={25}
                        name={field.name}
                        onBlur={field.handleBlur}
                        onChange={(e) => field.handleChange(e.target.value)}
                        placeholder='Jane Doe'
                        value={field.state.value}
                      />
                      {isInvalid && <FieldError errors={field.state.meta.errors} />}
                    </Field>
                  );
                }}
              </form.Field>
              <form.Field name='email'>
                {(field) => {
                  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;
                  return (
                    <Field data-invalid={isInvalid}>
                      <FieldLabel htmlFor={field.name}>Email</FieldLabel>
                      <Input
                        aria-invalid={isInvalid}
                        autoComplete='off'
                        id={field.name}
                        name={field.name}
                        onBlur={field.handleBlur}
                        onChange={(e) => field.handleChange(e.target.value)}
                        placeholder='m@example.com'
                        value={field.state.value}
                      />
                      {isInvalid && <FieldError errors={field.state.meta.errors} />}
                    </Field>
                  );
                }}
              </form.Field>
              <Field>
                <Field className='grid grid-cols-2 gap-4'>
                  <form.Field name='password'>
                    {(field) => {
                      const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;
                      return (
                        <Field data-invalid={isInvalid}>
                          <FieldLabel htmlFor={field.name}>Password</FieldLabel>
                          <Input
                            aria-invalid={isInvalid}
                            autoComplete='off'
                            id={field.name}
                            name={field.name}
                            onBlur={field.handleBlur}
                            onChange={(e) => field.handleChange(e.target.value)}
                            placeholder='Password'
                            value={field.state.value}
                          />
                          {isInvalid && <FieldError errors={field.state.meta.errors} />}
                        </Field>
                      );
                    }}
                  </form.Field>
                  <form.Field
                    name='confirmPassword'
                    validators={{
                      onChange: ({ value, fieldApi }) => {
                        if (value !== fieldApi.form.getFieldValue('password')) {
                          return {
                            message: 'Passwords do not match.',
                          };
                        }
                        return undefined;
                      },
                      onChangeListenTo: ['password'],
                    }}>
                    {(field) => {
                      const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;
                      return (
                        <Field data-invalid={isInvalid}>
                          <FieldLabel htmlFor={field.name}>Confirm Password</FieldLabel>
                          <Input
                            aria-invalid={isInvalid}
                            autoComplete='off'
                            id={field.name}
                            name={field.name}
                            onBlur={field.handleBlur}
                            onChange={(e) => field.handleChange(e.target.value)}
                            placeholder='Password Confirmation'
                            value={field.state.value}
                          />
                          {isInvalid && <FieldError errors={field.state.meta.errors} />}
                        </Field>
                      );
                    }}
                  </form.Field>
                </Field>
              </Field>
              <Field>
                <Button form='signup-form' type='submit'>
                  Create Account
                </Button>
              </Field>
              <FieldSeparator className='*:data-[slot=field-separator-content]:bg-card' />
              <FieldDescription className='text-center'>
                Already have an account? <Link to={`/authorize${window.location.search}`}>Sign in</Link>
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
        By clicking sign up, you agree to our <Link to='/terms'>Terms of Service</Link> and{' '}
        <Link to='/privacy'>Privacy Policy</Link>.
      </FieldDescription>
    </div>
  );
}
