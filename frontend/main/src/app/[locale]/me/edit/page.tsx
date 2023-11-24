'use client'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '$/components/ui/tabs'
import Privacy from './privacy'
import Profile from './profile'

const user = {
  name: 'dehwyy',
  email: 'dehwyy@google.com',
  customId: '',
  description: "Hello, I'm dehwyy and using Makoto but way longer sentence to test if it works.",
  dark_background: '#171717',
  light_background: '#2ccce7',
  image: '',

  isVerifiedEmail: false,
}

const Page = () => {
  return (
    <div className="min-h-screen mx-auto pt-28 w-[90%] md:w-2/3">
      <Tabs defaultValue="profile" className="w-full">
        <TabsList className="grid grid-cols-2">
          <TabsTrigger value="profile">User Profile</TabsTrigger>
          <TabsTrigger value="privacy">Privacy</TabsTrigger>
        </TabsList>

        {/* User profile  */}
        <TabsContent value="profile">
          <Profile
            name={user.name}
            email={user.email}
            customId={user.customId}
            description={user.description}
            dark_background={user.dark_background}
            light_background={user.light_background}
            image={user.image}
            isVerifiedEmail={user.isVerifiedEmail}
          />
        </TabsContent>

        {/* User privacy  */}
        <TabsContent value="privacy">
          <Privacy />
        </TabsContent>
      </Tabs>
    </div>
  )
}

export default Page
