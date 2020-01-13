metadata :name        => "user",
         :description => "Interact with users on a system",
         :author      => "exaldraen",
         :license     => "GPL-3.0-or-later",
         :version     => "0.4.0",
         :url         => "https://github.com/exaldraen/choria-plugin-user-agent",
         :provider    => "external",
         :timeout     => 15


action "list", :description => "List logged-in users on a system" do
  display :ok



  output :sessions,
         :description => "List of logged-in user sessions",
         :display_as  => "Session List",
         :type        => "array"

  summarize do
    summary(:sessions)
  end
end

action "kill", :description => "Kill all sessions belonging to the given user" do
  display :ok

  input :user,
        :prompt      => "Username",
        :description => "Name of the user whose sessions will be killed",
        :type        => :string,
        :validation  => :shellsafe,
        :maxlength   => 128,
        :optional    => false




end

