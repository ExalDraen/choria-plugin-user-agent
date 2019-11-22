metadata :name        => "user",
         :description => "Interact with users on a system",
         :author      => "Alexander Hermes (@ExalDraen)",
         :license     => "GNU General Public License v3.0",
         :version     => "0.1.0",
         :url         => "https://github.com/exaldraen/choria-plugin-user-agent",
         :provider    => "external",
         :timeout     => 15


action "echo", :description => "Echo text back" do
  display :ok

  input :message,
        :prompt      => "The message to send for echoing",
        :description => "This message will be echoed back",
        :type        => :string,
        :validation  => :shellsafe,
        :maxlength   => 512,
        :optional    => false




  output :message,
         :description => "Echo of the original message",
         :display_as  => "Echo",
         :type        => "string"

end

action "list", :description => "List logged-in users on a system" do
  display :ok



  output :user_list,
         :description => "List of logged-in users",
         :display_as  => "User LIst",
         :type        => "string"

end

