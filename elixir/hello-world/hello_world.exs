defmodule HelloWorld do
  @moduledoc """
  A module to say hello to the world.
  """

  @doc """
  Greets the user by name, or by saying "Hello, World!"
  if no name is given.
  """
  @spec hello(String.t()) :: String.t()
  def hello(name \\ "World") when is_binary(name), do: "Hello, #{name}!"
end
