"use client";
import React from "react";
import rhf, { useForm } from "react-hook-form";

interface FormData {
  firstName: string;
  lastName: string;
  countryCode: string;
  PhoneNumber: string;
  email: string;
  password: string;
}

const Input = ({
  register,
  errors,
  inputName,
}: {
  register: rhf.UseFormRegister<FormData>;
  errors: rhf.FieldErrors<FormData>;
  inputName: keyof FormData;
}) => {
  return (
    <div className="flex flex-col text-left gap-2 pb-4">
      <label>{inputName.charAt(0).toUpperCase() + inputName.slice(1)}: </label>
      <input
        type={
          inputName === "email"
            ? "email"
            : inputName === "password"
            ? "password"
            : "text"
        }
        {...register(inputName, {
          required: `${inputName} is required`,
          pattern: {
            value: /^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/,
            message: `enter a valid ${inputName}`,
          },
        })}
        className="border border-black rounded-lg px-4 py-2 w-full"
      />
      {errors.email && <p>{errors.email.message}</p>}
    </div>
  );
};

const Signup = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>();

  const onSubmit = (data: FormData) => {
    console.log(data);
    // Handle the form data here
  };

  return (
    <main className="flex items-center justify-center pt-32">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-[400px] max-w-lg border border-black rounded-lg px-8 py-4"
      >
        {(["firstName", "lastName", "countryCode", "PhoneNumber"] as const).map(
          (key) => (
            <Input register={register} errors={errors} inputName={key} />
          )
        )}

        <button
          type="submit"
          className="bg-green-600 px-4 py-2 rounded-lg text-white"
        >
          Sign Up
        </button>
      </form>
    </main>
  );
};

export default Signup;
